FROM public.ecr.aws/lambda/java:11.2023.06.12.17

COPY build/classes/java/main ${LAMBDA_TASK_ROOT}
CMD [ "org.hellojava.Handler" ]

