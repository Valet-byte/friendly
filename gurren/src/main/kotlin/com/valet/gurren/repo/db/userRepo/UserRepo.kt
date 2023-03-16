package com.valet.gurren.repo.db.userRepo

import com.valet.gurren.model.User
import jakarta.transaction.Transactional
import org.springframework.data.jpa.repository.Modifying
import org.springframework.data.jpa.repository.Query
import org.springframework.data.repository.CrudRepository
import org.springframework.data.repository.query.Param
import java.awt.print.Pageable

interface UserRepo : CrudRepository<User, String> {
    @Modifying
    @Query("UPDATE User u set u.isPremium = :status where u.username = :username")
    fun updateUserPremiumStatus(@Param("username") username: String, @Param("status") status: Boolean)

    @Modifying
    @Query("update User u set u.isEnable = false where u.username = :username")
    fun deleteUser(@Param("username") username: String)

    @Modifying
    @Query("update User u set u.isEnable = true where u.username = :username")
    fun recoveryUser(@Param("username") username: String)

    @Transactional
    @Modifying
    @Query(nativeQuery = true, value = "INSERT into friends (username1, username2) values (?, ?)")
    fun addFriend(username1: String, username2: String)

    @Query(nativeQuery = true, value = "SELECT username2 from friends WHERE username1 = :username AND not isempty(SELECT username1 from friends WHERE username2 = :username)")
    fun getFriends(@Param("username") username: String): List<String>

    @Query(nativeQuery = true, value = "SELECT username2 from friends WHERE username1 = :username AND isempty(SELECT username1 from friends WHERE username2 = :username)")
    fun getSubscribers(@Param("username") username : String): List<String>

    @Query(nativeQuery = true, value = "SELECT username1 from friends WHERE username2 = :username AND isempty(SELECT username2 from friends WHERE username1 = :username)")
    fun getFollowers(@Param("username") username : String): List<String>
}