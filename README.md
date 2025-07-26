# Conversor de Temperaturas - API

API para converter temperaturas entre Celsius, Fahrenheit e Kelvin.

---

## Tecnologias

- Go (Golang)  
- Framework Gin para API REST

---

## O que aprendi com este projeto

- Estrutura básica de pastas em projetos Go (organização com `controllers`, `services`, etc)  
- Como usar o framework Gin para criar APIs REST  
- Separação de responsabilidades entre camadas (controllers para rotas, services para lógica)  
- Tratamento de erros e validação de parâmetros  
- Formatação e resposta em JSON

---

## Funcionalidades

- Conversão de temperatura entre Celsius, Fahrenheit e Kelvin  
- Respostas em JSON  
- Validação de entrada com tratamento de erros

---

## Endpoints

- **GET** `https://tempconverter-api.onrender.com/converter/celsius?celsius=VALOR`  
  Converte de Celsius para Fahrenheit e Kelvin.  
  **Exemplo:** `https://tempconverter-api.onrender.com/converter/celsius?celsius=30`

- **GET** `https://tempconverter-api.onrender.com/converter/fahrenheit?fahrenheit=VALOR`  
  Converte de Fahrenheit para Celsius e Kelvin.  
  **Exemplo:** `https://tempconverter-api.onrender.com/converter/fahrenheit?fahrenheit=86`

- **GET** `https://tempconverter-api.onrender.com/converter/kelvin?kelvin=VALOR`  
  Converte de Kelvin para Celsius e Fahrenheit.  
  **Exemplo:** `https://tempconverter-api.onrender.com/converter/kelvin?kelvin=303.15`

---

## Resposta de exemplo

```json
{
  "celsius": "30.00",
  "fahrenheit": "86.00",
  "kelvin": "303.15"
}

``` 

## 🌐 Acesse os endpoints aqui

A API está publicada e disponível para testes no seguinte endereço:  
[https://tempconverter-api.onrender.com](https://tempconverter-api.onrender.com)

