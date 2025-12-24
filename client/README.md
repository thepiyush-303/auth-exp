# React + TypeScript + Vite

This template provides a minimal setup to get React working in Vite with HMR and some ESLint rules.

Currently, two official plugins are available:

- [@vitejs/plugin-react](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react) uses [Babel](https://babeljs.io/) (or [oxc](https://oxc.rs) when used in [rolldown-vite](https://vite.dev/guide/rolldown)) for Fast Refresh
- [@vitejs/plugin-react-swc](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react-swc) uses [SWC](https://swc.rs/) for Fast Refresh

## React Compiler

The React Compiler is not enabled on this template because of its impact on dev & build performances. To add it, see [this documentation](https://react.dev/learn/react-compiler/installation).

## Expanding the ESLint configuration

If you are developing a production application, we recommend updating the configuration to enable type-aware lint rules:

```js
export default defineConfig([
  # Auth Client (Register + Login)

  Minimal React (Vite + TS) frontend with two tabs: Register and Login. Errors and success messages from the backend are shown at the top of the form.

  ## Run

  1. Start the Go backend (from `go-auth`):

  ```bash
  go run .
  ```

  2. Start this client (from `go-auth/client`):

  ```bash
  npm install
  npm run dev
  ```

  Open the URL printed by Vite (usually http://localhost:5173).

  ## Notes

  - Requests are proxied to `http://localhost:3000` via Vite (`vite.config.ts`).
  - Register sends `name`, `email`, `password` as URL-encoded form fields to `/register`.
  - Login sends `email`, `password` as URL-encoded form fields to `/login`.
  - If `/login` is not implemented on the backend yet, the client will show the backend's error (e.g., 404) at the top.

