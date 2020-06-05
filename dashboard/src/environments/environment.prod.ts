var port = 9000;
declare var require: any

export const environment = {
  production: true,
  enqueuedUri: "http://localhost:"+port+"/enqueued",
  successfulUri:  "http://localhost:"+port+"/successful",
  scheduledUri: "http://localhost:"+port+"/scheduled",
  sumUri:"http://localhost:"+port+"sum",
  appVersion: require('../../package.json').version

};
