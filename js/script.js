// Copyright (c) 2023 Mohamad All rights reserved
//
// Created by: Mohamad
// Created on: May 2023
// This file contains the JS functions for index.html

/*
* This function will count up to a number of the user's choice
*/
function count() {
  // Get user input
  const numberGoal = parseInt(document.getElementById("number-goal").value)
  let answer = " "

  // Count
  for (let counter = 0; counter <= numberGoal; counter++) {
    // Output
    document.getElementById("answer").innerHTML = answer += counter + "<br>"
  }
}