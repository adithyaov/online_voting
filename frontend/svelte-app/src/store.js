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

import routeChangeHandeler from './handelers/routeChange.js';
store.on('routeChange', routeChangeHandeler);

import oppenBallotHandeler from './handelers/openBallot.js';
store.on('openBallot', oppenBallotHandeler);

export default store;
