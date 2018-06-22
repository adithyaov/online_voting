import { Store } from 'svelte/store.js';

const store = new Store({
  ballots: {
    codes: [],
    entities: {}
  },
  candidates: [],
  loading: false,
  registration: {
    value: '',
    submitted: false
  },
  route: 'home',
  notifications: []
});

const init = ctx => {
  store.on(ctx.eventName, ctx.handler);
};

const multiInit = ctxs => {
  ctxs.forEach(ctx => {
    init(ctx);
  });
};

import handlers from './handlers';
multiInit(handlers);

export default store;
