package com.valetbyte.gurren.model

data class Point(
    private val id: String,
    private val x: Double,
    private val y: Double,
    private val batteryCharge: UShort?,
    private val speed: UShort?,
    private val isUserInHome: Boolean?
)