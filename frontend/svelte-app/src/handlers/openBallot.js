function handler(eventData) {
  console.log(eventData.ballotCode);
}

const eventName = 'openBallot';

export default {
  handler: handler,
  eventName: eventName
};
