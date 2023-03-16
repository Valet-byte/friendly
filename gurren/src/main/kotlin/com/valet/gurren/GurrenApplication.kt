package com.valet.gurren

import com.google.auth.oauth2.GoogleCredentials
import com.google.firebase.FirebaseApp
import com.google.firebase.FirebaseOptions
import com.google.firebase.auth.FirebaseAuth
import com.google.firebase.database.FirebaseDatabase
import com.google.firebase.messaging.FirebaseMessaging
import org.springframework.beans.factory.annotation.Value
import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.context.annotation.Bean
import java.io.FileInputStream

@SpringBootApplication
class GurrenApplication {
    @Bean
    fun fireBaseApp(@Value("\${firebase.secret-key-path}") path: String): FirebaseApp {
        val serviceAccount = FileInputStream(path)

        val options = FirebaseOptions.builder()
            .setCredentials(GoogleCredentials.fromStream(serviceAccount))
            .build()

        return FirebaseApp.initializeApp(options)
    }

    @Bean
    fun firebaseMessaging(app: FirebaseApp) : FirebaseMessaging = FirebaseMessaging.getInstance(app)

    @Bean
    fun firebaseAuth(app: FirebaseApp) : FirebaseAuth = FirebaseAuth.getInstance(app)

    @Bean
    fun firebaseDatabase(app: FirebaseApp, @Value("\${firebase.database.url}") url: String) : FirebaseDatabase = FirebaseDatabase.getInstance(app, url)
}

fun main(args: Array<String>) {
    runApplication<GurrenApplication>(*args)
}
