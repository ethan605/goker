**Goker sequence diagrams**

- [Player flow](#player-flow)
- [Hand flow](#hand-flow)
- [Betting flow](#betting-flow)
- [Player folds](#player-folds)

# Player flow

```mermaid
sequenceDiagram
    participant Player
    participant Pot
    participant Hand
    participant Table

    Player->>+Table: Join table

    activate Table
    alt max Players exceeded
        Table-->>Player: Max players exceeded
    else
        Table-->>Player: Seat position
    end
    deactivate Table

    alt is first Player
        Table->>Table: Wait for more Players
    else 
        Table->>Table: Wait for next Hand
    end

    Table->>+Player: Start new hand
    Note over Player,Table: See "Hand flow"
    Player-->>-Table: Actions

    Player->>+Table: Left table
    Table->>-Table: Check and update dealer button
```

# Hand flow

```mermaid
sequenceDiagram
    participant Player
    participant Pot
    participant Hand
    participant Table

    Table->>Hand: Start new hand
    Table->>Pot: Start new hand

    Pot->>+Player: "Pre-flop" betting round
    activate Player
    alt is small blind
        Player-->>Pot: Bet small
    else is big blind
        Player-->>Pot: Bet big
    end
    deactivate Player
    
    Pot->>Hand: "Pre-flop" betting done

    loop until all hold cards are dealt
        Hand->>Player: Deal hold cards
    end

    Pot->>+Player: "Flop" betting round
    Note over Player,Table: See "Betting flow"
    Player-->>-Pot: Actions

    Pot->>+Player: "Turn" betting round
    Note over Player,Table: See "Betting flow"
    Player-->>-Pot: Actions

    Pot->>+Player: "River" betting round
    Note over Player,Table: See "Betting flow"
    Player-->>-Pot: Actions
```

# Betting flow

```mermaid
sequenceDiagram
    participant Player
    participant Pot
    participant Hand
    participant Table

    Pot->>Player: New betting round

    loop until active bets matched
    activate Player
        alt fold
            Note over Player,Hand: See "Player folds"
        else
            Player-->>Pot: Open/Call/Raise/Re-raise
        end
    deactivate Player
    end

    Pot->>Pot: Collect chips

    alt only 1 active bet left
        Pot->>Table: Hand ended
        Table->>Player: Win
        Table->>Pot: Reward chips
        Pot->>Player: Reward chips
        Table->>Hand: Clear hand
        Hand->>Player: Recall hold cards
    else
        opt all-ins
            Pot->>Pot: Calculate side pot
        end
        Pot->>Hand: Betting done
        Hand->>Hand: Burn cards
        Hand->>Hand: Deal community cards
        Hand->>Player: Community cards dealt
    end
```

## Player folds

```mermaid
sequenceDiagram
    participant Player
    participant Pot
    participant Hand
    participant Table

    Pot->>+Player: New betting round
    Player-->>-Pot: Fold
    opt
        Pot->>Pot: Collect chips
    end

    Pot->>Hand: Fold
    Hand->>Hand: Exclude player from next betting rounds
    Hand-->>+Player: Recall hold cards
    Player->>-Hand: Hold cards
```
