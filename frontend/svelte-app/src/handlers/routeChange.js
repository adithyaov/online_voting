function handler(eventData) {
  console.log(eventData.newRoute);
}

const eventName = 'routeChange';

export default {
  handler: handler,
  eventName: eventName
};
