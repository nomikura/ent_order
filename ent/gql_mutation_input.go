// Code generated by ent, DO NOT EDIT.

package ent

// CreateOrganizationInput represents a mutation input for creating organizations.
type CreateOrganizationInput struct {
	Name     string
	Priority int
	UserIDs  []int
}

// Mutate applies the CreateOrganizationInput on the OrganizationMutation builder.
func (i *CreateOrganizationInput) Mutate(m *OrganizationMutation) {
	m.SetName(i.Name)
	m.SetPriority(i.Priority)
	if v := i.UserIDs; len(v) > 0 {
		m.AddUserIDs(v...)
	}
}

// SetInput applies the change-set in the CreateOrganizationInput on the OrganizationCreate builder.
func (c *OrganizationCreate) SetInput(i CreateOrganizationInput) *OrganizationCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateOrganizationInput represents a mutation input for updating organizations.
type UpdateOrganizationInput struct {
	Name          *string
	Priority      *int
	ClearUsers    bool
	AddUserIDs    []int
	RemoveUserIDs []int
}

// Mutate applies the UpdateOrganizationInput on the OrganizationMutation builder.
func (i *UpdateOrganizationInput) Mutate(m *OrganizationMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Priority; v != nil {
		m.SetPriority(*v)
	}
	if i.ClearUsers {
		m.ClearUsers()
	}
	if v := i.AddUserIDs; len(v) > 0 {
		m.AddUserIDs(v...)
	}
	if v := i.RemoveUserIDs; len(v) > 0 {
		m.RemoveUserIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateOrganizationInput on the OrganizationUpdate builder.
func (c *OrganizationUpdate) SetInput(i UpdateOrganizationInput) *OrganizationUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateOrganizationInput on the OrganizationUpdateOne builder.
func (c *OrganizationUpdateOne) SetInput(i UpdateOrganizationInput) *OrganizationUpdateOne {
	i.Mutate(c.Mutation())
	return c
}