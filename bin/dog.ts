#!/usr/bin/env node
import 'source-map-support/register';
import * as cdk from 'aws-cdk-lib';
import { DogStack } from '../lib/dog-stack';

const app = new cdk.App();
new DogStack(app, 'DogStack', {
  env: {
    region: 'us-west-2',
  }
});