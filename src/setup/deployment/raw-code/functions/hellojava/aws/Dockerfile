FROM gradle:8.1.1-jdk11 AS build
COPY --chown=gradle:gradle . /home/gradle/src
WORKDIR /home/gradle/src
RUN gradle build --no-daemon 

# Copy the function artifact and dependencies onto a clean base
FROM public.ecr.aws/lambda/java:11.2023.06.12.17

COPY --from=build /home/gradle/src/build/classes/java/main ${LAMBDA_TASK_ROOT}
COPY --from=build /home/gradle/src/build/dependency/* ${LAMBDA_TASK_ROOT}/lib/

CMD [ "org.hellojava.Handler" ]
