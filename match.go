package matchmaking

func match(job *procesJob) {

	switch job.req.Command {
	case "FIND":
		matchFind(job)
	case "STATUS":
	case "QUIT":
	case "JOIN":

	}

}
