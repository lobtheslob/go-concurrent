package PerformAsync

// This function performs the given task concurrently by spawing a goroutine


func performAsyncTasks(task []Task) {
	for _, task := range tasks {
	  // This will spawn a separate goroutine to carry out this task.
	  // This call is non-blocking
	  go task.Execute()
	}
  }