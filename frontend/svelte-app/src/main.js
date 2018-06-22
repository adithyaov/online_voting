import App from './App.html';
import 'normalize.css/normalize.css';
import 'spectre.css/dist/spectre.min.css';
import './assets/global.css';
import store from './store.js';

const app = new App({
  target: document.body,
  store
});

window.onhashchange = () =>
  store.fire('routeChange', { nowRoute: location.hash });

export default app;
