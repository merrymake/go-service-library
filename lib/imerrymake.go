package lib

type IMerrymake interface {
	// Used to link actions in the Merrymake.json file to code.
	// # Arguments
	// * `action` -- the action from the Merrymake.json file
	// * `handler` -- the code to execute when the action is triggered
	// # Returns
	// The Merrymake builder to define further actions
	Handle(string, func([]byte, Envelope)) IMerrymake
	// Used to define code to run after deployment but before release. Useful for smoke tests or database consolidation. Similar to an 'init container'
	// # Arguments
	// * `handler` -- the code to execute
	Initialize(func())
}
