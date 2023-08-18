func solution(players []string, callings []string) []string {
	playerIndex := make(map[string]int)
	for i, player := range players {
		playerIndex[player] = i
	}

	sort.Slice(players, func(i, j int) bool {
		return playerIndex[players[i]] < playerIndex[players[j]]
	})

	for _, calling := range callings {
		if i, ok := playerIndex[calling]; ok && i > 0 {
			players[i], players[i-1] = players[i-1], players[i]
			playerIndex[players[i]] = i
			playerIndex[players[i-1]] = i - 1
		}
	}

	return players
}