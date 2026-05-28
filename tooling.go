package main

func delete_at_index(slice []GameObject, index int) []GameObject {
	return append(slice[:index], slice[index+1:]...)
}
