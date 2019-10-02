// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_database "github.com/oracle/oci-go-sdk/database"
)

func DatabaseBackupDestinationDataSource() *schema.Resource {
	fieldMap := make(map[string]*schema.Schema)
	fieldMap["backup_destination_id"] = &schema.Schema{
		Type:     schema.TypeString,
		Required: true,
	}
	return GetSingularDataSourceItemSchema(DatabaseBackupDestinationResource(), fieldMap, readSingularDatabaseBackupDestination)
}

func readSingularDatabaseBackupDestination(d *schema.ResourceData, m interface{}) error {
	sync := &DatabaseBackupDestinationDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).databaseClient

	return ReadResource(sync)
}

type DatabaseBackupDestinationDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_database.DatabaseClient
	Res    *oci_database.GetBackupDestinationResponse
}

func (s *DatabaseBackupDestinationDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *DatabaseBackupDestinationDataSourceCrud) Get() error {
	request := oci_database.GetBackupDestinationRequest{}

	if backupDestinationId, ok := s.D.GetOkExists("backup_destination_id"); ok {
		tmp := backupDestinationId.(string)
		request.BackupDestinationId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "database")

	response, err := s.Client.GetBackupDestination(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *DatabaseBackupDestinationDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	associatedDatabases := []interface{}{}
	for _, item := range s.Res.AssociatedDatabases {
		associatedDatabases = append(associatedDatabases, AssociatedDatabaseDetailsToMap(item))
	}
	s.D.Set("associated_databases", associatedDatabases)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.ConnectionString != nil {
		s.D.Set("connection_string", *s.Res.ConnectionString)
	}

	if s.Res.DefinedTags != nil {
		s.D.Set("defined_tags", definedTagsToMap(s.Res.DefinedTags))
	}

	if s.Res.DisplayName != nil {
		s.D.Set("display_name", *s.Res.DisplayName)
	}

	s.D.Set("freeform_tags", s.Res.FreeformTags)

	if s.Res.LifecycleDetails != nil {
		s.D.Set("lifecycle_details", *s.Res.LifecycleDetails)
	}

	if s.Res.LocalMountPointPath != nil {
		s.D.Set("local_mount_point_path", *s.Res.LocalMountPointPath)
	}

	s.D.Set("state", s.Res.LifecycleState)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	s.D.Set("type", s.Res.Type)

	s.D.Set("vpc_users", s.Res.VpcUsers)

	return nil
}