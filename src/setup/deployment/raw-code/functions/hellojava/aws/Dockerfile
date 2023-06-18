FROM public.ecr.aws/lambda/java:11.2023.06.12.17

COPY build/classes/java/main ${LAMBDA_TASK_ROOT}
COPY build/dependency/* ${LAMBDA_TASK_ROOT}/lib/
CMD [ "org.hellojava.Handler" ]

