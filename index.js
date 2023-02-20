// import the aws sdk
const AWS = require('aws-sdk');

module.exports.handler = async (event) => {

  // return cost and usage data for the last 30 days
  const costExplorer = new AWS.CostExplorer();
  const params = {
    Granularity: 'DAILY',
    TimePeriod: {
      Start: '2022-06-01',
      End: '2023-01-01'
    },
    Metrics: [
      'UnblendedCost',
    ],
    GroupBy: [
      {
        Type: 'DIMENSION',
        Key: 'SERVICE'
      },
    ],
  };

  const data = await costExplorer.getCostAndUsage(params).promise();

  // return the data
  return {
    statusCode: 200,
    data
  };
};
