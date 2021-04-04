**Table of content**

- [Player joins and leaves table](#player-joins-and-leaves-table)
- [Hand interactions](#hand-interactions)
- [Normal betting rounds](#normal-betting-rounds)
- [Player folds](#player-folds)

# Player joins and leaves table

```mermaid
sequenceDiagram
    participant Player
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
    Player-->>-Table: Join new hand

    Player->>+Table: Left table
    Table->>-Table: Check and update dealer button
```

# Hand interactions

```mermaid
sequenceDiagram
    participant Player
    participant Pot
    participant Hand
    participant Table

    Table->>Hand: Start new hand
    Table->>Pot: Start new hand

    Pot->>+Player: Pre-flop betting round
    activate Player
    alt is small blind
        Player-->>Pot: Bet small
    else is big blind
        Player-->>Pot: Bet big
    end
    deactivate Player
    
    Pot->>Hand: Pre-flop betting done

    loop until all hold cards are dealt
        Hand->>Player: Deal hold cards
    end

    Pot->>+Player: Extra betting needed
    Note over Player,Table: See "Normal betting rounds"
    Player-->>-Pot: Actions

    Pot->>+Player: Turn betting round
    Note over Player,Table: See "Normal betting rounds"
    Player-->>-Pot: Actions

    Pot->>+Player: River betting round
    Note over Player,Table: See "Normal betting rounds"
    Player-->>-Pot: Actions
```

# Normal betting rounds

```mermaid
sequenceDiagram
    participant Player
    participant Pot
    participant Hand
    participant Table

    loop until active bets matched
    Pot->>Player: New betting round

    activate Player
        alt fold
            Note over Player,Hand: See "Player folds"
        else
            Player-->>Pot: Call/Raise
        end
    deactivate Player
    end

    Pot->>Pot: Collect chips

    alt only 1 active bet left
        Pot->>Table: Hand ended
        Table->>Player: Win
        Pot->>Player: Reward chips
        Table->>Hand: Clear hand
        Hand->>Player: Recall hold cards
    else
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

    Pot->>+Player: Betting round
    Player-->>-Pot: Fold
    opt
        Pot->>Pot: Collect chips
    end

    Player->>+Hand: Fold
    Hand->>Hand: Exclude from next betting round
    Hand-->>-Player: Recall hold cards
```
