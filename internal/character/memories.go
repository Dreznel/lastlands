package character

import "fmt"

type CollectedMemory struct {
	Memory       string
	Details      string
	DiscoveredAt string
}

type CoreAttributes struct {
	First, Second, Third CoreAttribute
}

type CoreAttribute interface {
	GetDescription() string
	UpdateDescription(newDescription string)
	IsLost() bool
	SetIsLost(b bool)
	GetType() string
}

type CoreMemory struct {
	IsMemoryLost bool
	Description  string
}

func (cm *CoreMemory) GetDescription() string {
	return cm.Description
}

func (cm *CoreMemory) UpdateDescription(newDescription string) {
	cm.Description = newDescription
}

func (cm *CoreMemory) IsLost() bool {
	return cm.IsMemoryLost
}
func (cm *CoreMemory) SetIsLost(b bool) {
	cm.IsMemoryLost = b
}
func (cm *CoreMemory) GetType() string {
	return "Core Memory"
}

type CorePurpose struct {
	Description string
}

func (cp *CorePurpose) GetDescription() string {
	return cp.Description
}

func (cp *CorePurpose) UpdateDescription(newDescription string) {
	cp.Description = newDescription
}

func (cp *CorePurpose) IsLost() bool {
	return false // Core Purposes cannot be Lost
}
func (cp *CorePurpose) SetIsLost(b bool) {
	fmt.Println("Core Purposes cannot be lost") // Core Purposes cannot be lost
}
func (cp *CorePurpose) GetType() string {
	return "Core Purpose"
}
