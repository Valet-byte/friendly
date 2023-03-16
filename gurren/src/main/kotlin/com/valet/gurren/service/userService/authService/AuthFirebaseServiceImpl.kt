package com.valet.gurren.service.userService.authService

import com.google.firebase.auth.FirebaseAuth
import com.valet.gurren.exception.UserAlreadyExists
import com.valet.gurren.model.User
import com.valet.gurren.model.UserData
import com.valet.gurren.repo.cache.userCache.UserCache
import com.valet.gurren.repo.db.userRepo.UserRepo
import org.springframework.stereotype.Service
import java.lang.IllegalArgumentException

@Service
class AuthFirebaseServiceImpl(private val auth: FirebaseAuth,
                              private val cache: UserCache,
                              private val userRepo: UserRepo) : AuthService {
    override fun auth(userId: String): User {
        return if (cache.hasUser(userId)){
            cache.getUser(userId)
        } else {
            val user = userRepo.findById(auth.getUser(userId).displayName).get()
            cache.save(user, userId)
            user
        }
    }

    override fun registration(userId: String, userData: UserData) {
        val username = auth.getUser(userId).displayName
        if (username != null){
            var user = User.createUser(userData)
            user.username = username
            try {
                user = userRepo.save(user)
            } catch (e : Exception) {
                throw UserAlreadyExists("User with name : $username already exists!" )
            }

            cache.save(user, userId)
        } else {
            throw IllegalArgumentException("Not found user with id : $userId !")
        }
    }




}