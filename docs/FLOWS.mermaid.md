# Goker flows

- [Player flow](#player-flow)
- [Dealer flow](#dealer-flow)
- [Betting flow](#betting-flow)
- [Player folds](#player-folds)
- [Winner flow](#winner-flow)

## Player flow

```mermaid
sequenceDiagram
    participant Player
    participant Pot
    participant Dealer
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
        Table->>Table: Wait for next Dealer
    end

    Table->>Dealer: Start new hand
    Note over Player,Table: See "Dealer flow"

    Player->>+Table: Leave table
    Table->>-Table: Check and update dealer button
```

## Dealer flow

```mermaid
sequenceDiagram
    participant Player
    participant Pot
    participant Dealer
    participant Table

    Table->>Dealer: Start new hand
    Dealer->>Pot: New betting rounds

    Pot->>+Player: "Pre-flop" betting round
    activate Player
    alt is small blind
        Player-->>Pot: Bet small
    else is big blind
        Player-->>Pot: Bet big
    end
    deactivate Player
    
    Pot->>Dealer: "Pre-flop" betting done

    loop until all hold cards are dealt
        Dealer->>Player: Deal hold cards
    end

    Pot->>Player: "Flop" betting round
    activate Player
    Note over Player,Table: See "Betting flow"

    Pot->>Player: "Turn" betting round
    Note over Player,Table: See "Betting flow"

    Pot->>Player: "River" betting round
    Note over Player,Table: See "Betting flow"
    deactivate Player

    Dealer->>+Player: Showdown
    Player-->>-Dealer: Face-up cards
    Dealer->>Dealer: Rank player hands
    Note over Player,Table: See "Winner flow"
```

## Betting flow

```mermaid
sequenceDiagram
    participant Player
    participant Pot
    participant Dealer
    participant Table

    Pot->>Player: New betting round

    loop until active bets matched
    activate Player
        alt fold
            Note over Player,Dealer: See "Player folds"
        else
            Player-->>Pot: Open/Call/Raise/Re-raise
        end
    deactivate Player
    end

    Pot->>Pot: Collect chips

    alt only 1 active bet left
        Pot->>Dealer: Hand ended
        Note over Player,Table: See "Winner flow"
    else
        opt all-ins
            Pot->>Pot: Calculate side pot
        end
        Pot->>Dealer: Betting done
        Dealer->>Dealer: Burn cards
        Dealer->>Dealer: Deal community cards
        Dealer->>Player: Community cards dealt
    end
```

## Player folds

```mermaid
sequenceDiagram
    participant Player
    participant Pot
    participant Dealer
    participant Table

    Pot->>+Player: New betting round
    Player-->>-Pot: Fold

    Pot->>Dealer: Fold
    Dealer->>Dealer: Exclude player from next betting rounds
    Dealer->>+Player: Recall hold cards
    Player-->>-Dealer: Hold cards
```

## Winner flow

```mermaid
sequenceDiagram
    participant Player
    participant Pot
    participant Dealer
    participant Table

    Dealer->>Player: Winner
    Dealer->>Pot: Reward chips
    Pot->>Player: Reward chips
    Dealer->>+Player: Recall hold cards
    Player-->>-Dealer: Hold cards
    Dealer->>Table: Hand ended
```
