package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nimbo1999/temperature-challenge/internal/infra/observability"
	"github.com/nimbo1999/temperature-challenge/internal/infra/web/dto"
	"github.com/nimbo1999/temperature-challenge/internal/repository"
	"github.com/nimbo1999/temperature-challenge/internal/services"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	traceIdHeader := r.Header.Get(observability.TraceIdHeader)
	traceId, err := trace.TraceIDFromHex(traceIdHeader)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error() + "\n"))
		return
	}
	ctx := r.Context()
	ctx = trace.ContextWithSpanContext(ctx, trace.NewSpanContext(trace.SpanContextConfig{
		TraceID: traceId,
	}))

	// Configure span for handler
	tr := otel.GetTracerProvider().Tracer("weather-handler-component")
	_, span := tr.Start(ctx, "weather-span")
	span.SpanContext()
	defer span.End()

	cep := chi.URLParam(r, "cep")

	span.SetAttributes(attribute.String("CEP", cep))

	viacepService := services.NewViaCepService(repository.NewCepRepository())
	address, err := viacepService.GetData(cep)

	switch err {
	case services.ErrCepNotValid:
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(err.Error() + "\n"))
		return
	case services.ErrCepNotFound:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error() + "\n"))
		return
	case nil:
		break
	default:
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error() + "\n"))
		return
	}

	weatherService := services.NewWeatherService(repository.NewWeatherRepository())
	temperature, err := weatherService.GetData(*address)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error() + "\n"))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto.TemperatureResponseDTO{
		Celsius:    temperature.Celsius,
		Fahrenheit: temperature.Fahrenheit,
		Kelvin:     temperature.Kelvin,
	})
}
