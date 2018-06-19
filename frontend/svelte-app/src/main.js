import App from './App.html';
import 'normalize.css/normalize.css';
import 'spectre.css/dist/spectre.min.css';
import './assets/global.css';

const app = new App({
  target: document.body,
  data: {
    name: 'world'
  }
});

export default app;
