import org.jetbrains.kotlin.gradle.tasks.KotlinCompile

@Suppress("DSL_SCOPE_VIOLATION")
plugins {
    alias(libs.plugins.springframework.boot)
    alias(libs.plugins.spring.dependencies.manager)
    alias(libs.plugins.kotlin.jvm)
    alias(libs.plugins.kotlin.serialization)
    alias(libs.plugins.kotlin.spring)
}

group = "com.ams"
version = "0.0.1-SNAPSHOT"
java.sourceCompatibility = JavaVersion.VERSION_17

repositories {
    mavenCentral()
    gradlePluginPortal()
}

dependencies {
    // jjwt
    implementation(libs.jjwt.api)
    implementation(libs.jjwt.gson)
    implementation(libs.jjwt.impl)

    // Kotlinx coroutine
    implementation(libs.kotlinx.coroutines.core)
    implementation(libs.kotlinx.coroutines.reactor)

    // Kotlin reflect
    implementation(libs.kotlin.reflect)

    // Kotlinx serialization
    implementation(libs.kotlinx.serialization.json)

    // Spring
    implementation(libs.spring.boot.starter.web)
    implementation(libs.spring.boot.starter.data.r2dbc)
    implementation(libs.spring.boot.starter.security)
    implementation(libs.spring.data.relational)

    // Spring security
    implementation(libs.spring.security.core)


    // R2dbc mysql driver
    runtimeOnly(libs.jasync.r2dbc.mysql)
//    developmentOnly("org.springframework.boot:spring-boot-devtools")
    annotationProcessor(libs.spring.boot.configuration.processor)
//    implementation("org.springframework.boot:spring-boot-starter")
    testImplementation(libs.spring.boot.starter.test)
}

tasks.withType<KotlinCompile> {
    kotlinOptions {
        freeCompilerArgs = listOf("-Xjsr305=strict", "-opt-in=kotlin.io.encoding.ExperimentalEncodingApi")
        jvmTarget = "17"
    }
}

tasks.withType<Test> {
    useJUnitPlatform()
}
