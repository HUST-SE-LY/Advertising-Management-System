package com.ams.backend

import com.ams.backend.router.router
import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.context.support.beans

@SpringBootApplication(scanBasePackages = ["com.ams.backend"])
class BackendApplication

val beans = beans {
    bean(::router)
}


fun main(args: Array<String>) {
    runApplication<BackendApplication>(*args) {
        addInitializers(beans)
    }
}
