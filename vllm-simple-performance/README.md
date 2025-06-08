## vLLM Performance App

### Description
- golang app that sends a series of prompts to a vLLM endpoint (`/v1/completions`) to check performance
- polls and downloads metrics from the endpoint simultaneously
- intended to validate vLLM tuning parameters

### Usage
1. set the `config.yaml` (example)[./config.yaml] in the directory with 
  - url for the vLLM endpoint
  - url for the vLLM metrics endpoint
2. run the app `./performance`
3. collect the output metrics files