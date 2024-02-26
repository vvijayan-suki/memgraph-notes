package label

const (
	// Note Top level node in the graph
	Note = "Note"
	// NMSEvent c-line node
	NMSEvent = "NMS_Event"
	// Section label to identify Section node
	Section = "Section"
	// Metadata label to identify Metadata node
	Metadata = "Metadata"
	// MetadataEntry label to identify metadata entry received from create note event
	MetadataEntry = "Metadata_Entry"
	// SectionContent label to identify section content received from create note event
	SectionContent = "Section_Content"
	// SectionS2Entry label to identify section s2 entry node
	SectionS2Entry = "Section_S2_Entry"
	// MacroContent label to identify macro content extracted from macro event
	MacroContent = "Macro_Content"
	// VersionedComposition label to identify versioned composition for existing notes
	VersionedComposition = "Versioned_Composition"
)

// Relationship Labels
const (
	// Has label, relationship between Note and Section & Metadata node
	Has = "has"

	// CNext label, relationship used to identify c-line
	CNext = "c_next"
	// HNext label, relationship used to identify h-line
	HNext = "h_next"

	// Added label, relationship connecting and event node which adds new nodes
	// eg: create note event, adds Section, Metadata, MetadataEntry and SectionContent node
	Added = "added"
)
