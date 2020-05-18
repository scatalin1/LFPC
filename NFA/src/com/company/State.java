package com.company;

import java.util.ArrayList;

public class State {
    String name;
    ArrayList<Transition> transitions;

    State(String name){

        this.name = name;
        transitions = new ArrayList<Transition>();

    }

    public void addTransition(Transition t){

        transitions.add(t);

    }
}
