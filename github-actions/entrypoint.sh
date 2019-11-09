#!/bin/bash

set -ue

eval "$(/awscredswrap --role-arn ${INPUT_ROLE_ARN} --role-session-name ${INPUT_ROLE_SESSION_NAME} --duration-seconds ${INPUT_DURATION_SECONDS})"

echo ::add-mask::${AWS_ACCESS_KEY_ID}
echo ::add-mask::${AWS_SECRET_ACCESS_KEY}
echo ::add-mask::${AWS_SESSION_TOKEN}
echo ::add-mask::${AWS_DEFAULT_REGION}

echo ::set-env name=AWS_ACCESS_KEY_ID::${AWS_ACCESS_KEY_ID}
echo ::set-env name=AWS_SECRET_ACCESS_KEY::${AWS_SECRET_ACCESS_KEY}
echo ::set-env name=AWS_SESSION_TOKEN::${AWS_SESSION_TOKEN}
echo ::set-env name=AWS_DEFAULT_REGION::${AWS_DEFAULT_REGION}
