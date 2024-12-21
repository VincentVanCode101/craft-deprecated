FROM maven:3.9.6-eclipse-temurin-21 AS builder

ARG UID=1000
ARG GID=1000
ARG ARTIFACT_ID=default-project-name

WORKDIR /build-space


RUN mvn io.quarkus.platform:quarkus-maven-plugin:3.17.5:create \
    -DprojectGroupId=org.acme \
    -DprojectArtifactId=${ARTIFACT_ID} \
    -Dextensions='rest'

RUN chown -R 1000:1000 /build-space

CMD ["tail", "-f", "/dev/null"]