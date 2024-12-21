FROM maven:3.9.6-eclipse-temurin-21-jammy AS dev

WORKDIR /workspace

COPY ./pom.xml ./

RUN mvn dependency:go-offline

COPY src/ ./src/
