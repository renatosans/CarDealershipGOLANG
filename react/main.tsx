import App from './App'
import React from 'react'
import ReactDOM from 'react-dom/client'
import Layout from './components/Layout'

// O  TAURI é utilizado para aplicações semelhantes ao Electron porem não foi utilizado aqui
// com este propósito, ele está sendo utilizado para prototipagem no estilo monolito, algo
// semelhande ao que o NEXTjs, Nuxt ou REMIX fornecem
//
// Este projeto não usa o Custom Protocol fornecido pelo TAURI, pois não se enquadra nas práticas
// mais comuns de mercado  REST API, GraphQL, tRPC, etc...
// Ao inves disso é utilizada uma API padrão de mercado (REST API)

// The purpose of this project is to provide a developer experience similar to NEXTjs, Nuxt
// or REMIX, using a frontend and backend packed toghether
// The protocol provided by TAURI was not used in this project, instead a REST API is used

// Frontend at http://127.0.0.1:1420/
// Backend at http://127.0.0.1:8080/api/cars


ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <Layout>
      <App />
    </Layout>
  </React.StrictMode>
)
