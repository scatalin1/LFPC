package com.company;

public class Transition {
    String letter;
    State next;

    Transition(String letter, State next){

        this.letter = letter;
        this.next = next;

    }

}

