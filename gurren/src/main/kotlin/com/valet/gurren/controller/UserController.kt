package com.valet.gurren.controller

import com.valet.gurren.model.Description
import com.valet.gurren.model.User
import com.valet.gurren.model.UserData
import com.valet.gurren.model.UserProfileSetting
import org.springframework.http.ResponseEntity
import org.springframework.web.bind.annotation.DeleteMapping
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.PathVariable
import org.springframework.web.bind.annotation.PostMapping
import org.springframework.web.bind.annotation.RequestBody
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RequestParam
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("/user")
class UserController {

    @GetMapping("/{username}")
    fun getUser(@PathVariable("username") username: String) : User{

    }

    @GetMapping("/description/{username}")
    fun getUserDescription(@PathVariable("username") username: String) : Description{

    }

    @PostMapping("/")
    fun createUser(@RequestParam("user_id") userId: String,
                   @RequestParam("session_id") sessionId : String,
                   @RequestBody userData: UserData) : User{

    }

    @GetMapping("/friends")
    fun getFriends(@RequestParam("user_id") userId: String,
                   @RequestParam page: Int,
                   @RequestParam size: Int
    ) : List<User> {

    }

    @PostMapping("/friend/{username}")
    fun addFriend(@RequestParam("user_id") userId: String,
                  @PathVariable("username") username: String): ResponseEntity<String>{

    }

    @DeleteMapping("/friend/{username}")
    fun deleteUser(@RequestParam("user_id") userId: String,
                  @PathVariable("username") username: String): ResponseEntity<String>{

    }

    @DeleteMapping("/")
    fun deleteUser(@RequestParam("user_id") userId: String): ResponseEntity<String>{

    }

    @PostMapping("/settings")
    fun updateSettings(@RequestParam("user_id") userId: String,
                       @RequestBody userProfileSetting: UserProfileSetting) : ResponseEntity<String>{

    }

    @GetMapping("/settings")
    fun getUserSettings(@RequestParam("user_id") userId: String) : UserProfileSetting{

    }
}