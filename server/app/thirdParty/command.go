package thirdParty

import (
	"fantasytipster/server/app/expertTeams"
	"fantasytipster/server/app/experts"
	"fantasytipster/server/app/services/skapp"
	"fantasytipster/server/app/users"
	"stolencarsproject/server/internal/command"
)

func SetThirdPartyCommands() {
	command := command.Command{}

	// command = getTournamentCommand()
	// command.AddCommandWithArgs(insertTournamentData)

	// command = getTeamCommand()
	// command.AddCommandWithArgs(insertTeamData)

	// command = getMatchResultCommand()
	// command.AddCommandWithArgs(insertMatchResultData)

	// command = generateReferralCodeCommand()
	// command.AddCommandWithArgs(generateReferralCodeForUsers)

	// command = generateU11WinningsCommand()
	// command.AddCommandWithArgs(genrateU11WinningsForExpertTeams)

	command = generateNPoliceOfficersCommand()
	command.AddCommandWithArgs(generateNPoliceOfficersProfiles)
}

func getTournamentCommand() command.Command {
	return command.Command{
		Name:        "insertTournamentsData",
		Description: "add tournaments data to database",
		Category:    "Third Party Commands",
	}
}

func insertTournamentData(args ...string) {
	skapp.InsertTournamentDataCommandAction()
}

func getTeamCommand() command.Command {
	return command.Command{
		Name:        "insertTeamsData",
		Description: "add tournaments data to database",
		Category:    "Third Party Commands",
	}
}

func insertTeamData(args ...string) {
	team := args[0]
	skapp.InsertTeamData(team)
}

func getMatchResultCommand() command.Command {
	return command.Command{
		Name:        "insertMatchResultsData",
		Description: "add tournaments data to database",
		Category:    "Third Party Commands",
	}
}

func insertMatchResultData(args ...string) {
	matchSlug := args[0]
	skapp.InsertMatchResultData(matchSlug)
	expertTeams.UpdateMatchScorePerExpertTeamForAMatch(matchSlug)
	expertTeams.UpdateMatchPercentilePerExpertTeamForAMatch(matchSlug)
	experts.UpdateExpertScoreRankPercentileAndMatchesPredicted(matchSlug)
}

func generateReferralCodeCommand() command.Command {
	return command.Command{
		Name:        "generateReferralCodeForUsers",
		Description: "add referral codes for the users whose referral code not generated",
		Category:    "Third Party Commands",
	}
}

func generateU11WinningsCommand() command.Command {
	return command.Command{
		Name:        "generateU11WinningsForExpertTeams",
		Description: "add U11 winnings for expert teams to database",
		Category:    "Third Party Commands",
	}
}

func generateReferralCodeForUsers(args ...string) {
	users.GenerateReferralCodeForUsers()
}

func genrateU11WinningsForExpertTeams(args ...string) {
	expertTeams.UpdateU11WinningsForAllExpertTeams()
}

func generateNPoliceOfficersCommand() command.Command {
	return command.Command{
		Name:        "generateNPoliceOfficersProfiles",
		Description: "Add Win Ratio and H2H Ratio to database",
		Category:    "Third Party Commands",
	}
}

func generateNPoliceOfficersProfiles(args ...string) {
	experts.UpdateWinRatioAndH2HRatioForExpertProfiles()
}
