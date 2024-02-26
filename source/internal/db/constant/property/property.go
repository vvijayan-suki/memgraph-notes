package property

const (
	// ID primary id of node
	// eg: For Section node, ID is the section_id
	// eg: For an event node, ID is the event_id
	ID = "id"

	// SecondaryID unique id for a node, ideal for searching in graph
	// Eg: In the graph database, for different notes, Section nodes can be created with same section_id
	// Searching for a node with id=<section_id> would return multiple nodes from different notes
	// secondary_id is used to uniquely identify a node
	SecondaryID = "secondary_id"

	// Proto property, for a event node (on c-line) to store nmsEvent in bytes format
	Proto = "proto"

	// Name property, for a Section Node to hold the section name
	Name = "name"

	// StartCursorPosition property, for a Section Node to hold the starting cursor position
	StartCursorPosition = "start_cursor_position"
	// EndCursorPosition property, for a Section Node to hold the ending cursor position
	EndCursorPosition = "end_cursor_position"

	// Location property, for deciding whether the cursor is in the title or the content
	Location = "location"

	// Deleted property for a section which has been removed by undo
	Deleted = "deleted"

	// Status property for storing the note status
	Status = "status"
)
