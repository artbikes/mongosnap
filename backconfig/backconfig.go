package backconfig

import (
	"io/ioutil"

	"github.com/crielly/mongosnap/logger"
	"github.com/ghodss/yaml"
)

// Config describes the configuration of a MongoD process
type Config struct {
	Cluster struct {

		Storage struct {
			VolumeGroup 	string `json:"volumeGroup"`
			LogicalVolume 	string `json:"logicalVolume"`
			FileSystem 		string `json:"fileSystem"`
		} `json:"storage"`

		Snapshot struct {
			MountPath 		string `json:"mountPath"`
			Opts 			string `json:"opts"`
			SnapshotName 	string `json:"snapshotName"`
			Size 			string `json:"size"`
		} `json:"snapshot"`

		ReplicaConfs []string `json:"replicaConfs"`
		
	} `json:"cluster"`

	S3 struct {
		Bucket string `json:"bucket"`
		ObjectPath string `json:"objectPath"`
	} `json:"s3"`
}

// BackConfig prints a test MongoD Replica Config
func BackConfig(configPath string) (replconf Config, err error) {

	y, err := ioutil.ReadFile(configPath)
	if err != nil {
		logger.Error.Println(err)
	}

	err = yaml.Unmarshal(y, &replconf)
	if err != nil {
		logger.Error.Println(err)
	}

	return replconf, err
}
