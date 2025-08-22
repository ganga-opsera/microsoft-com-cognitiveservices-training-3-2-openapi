package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// BoundingBox represents the BoundingBox schema from the OpenAPI specification
type BoundingBox struct {
	Height float32 `json:"height"` // Height.
	Left float32 `json:"left"` // Coordinate of the left boundary.
	Top float32 `json:"top"` // Coordinate of the top boundary.
	Width float32 `json:"width"` // Width.
}

// ImageUrl represents the ImageUrl schema from the OpenAPI specification
type ImageUrl struct {
	Url string `json:"url"` // Url of the image.
}

// Export represents the Export schema from the OpenAPI specification
type Export struct {
	Downloaduri string `json:"downloadUri,omitempty"` // URI used to download the model.
	Flavor string `json:"flavor,omitempty"` // Flavor of the export. These are specializations of the export platform. Docker platform has valid flavors: Linux, Windows, ARM. Tensorflow platform has valid flavors: TensorFlowNormal, TensorFlowLite. ONNX platform has valid flavors: ONNX10, ONNX12.
	Newerversionavailable bool `json:"newerVersionAvailable,omitempty"` // Indicates an updated version of the export package is available and should be re-exported for the latest changes.
	Platform string `json:"platform,omitempty"` // Platform of the export.
	Status string `json:"status,omitempty"` // Status of the export.
}

// SuggestedTagAndRegionQuery represents the SuggestedTagAndRegionQuery schema from the OpenAPI specification
type SuggestedTagAndRegionQuery struct {
	Results []StoredSuggestedTagAndRegion `json:"results,omitempty"` // Result of a suggested tags and regions request of the untagged image.
	Token SuggestedTagAndRegionQueryToken `json:"token,omitempty"` // Contains properties we need to fetch suggested tags for. For the first call, Session and continuation set to null. Then on subsequent calls, uses the session/continuation from the previous SuggestedTagAndRegionQuery result to fetch additional results.
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	TypeField string `json:"type"` // Gets or sets the type of the tag.
	Description string `json:"description"` // Gets or sets the description of the tag.
	Id string `json:"id,omitempty"` // Gets the Tag ID.
	Imagecount int `json:"imageCount,omitempty"` // Gets the number of images with this tag.
	Name string `json:"name"` // Gets or sets the name of the tag.
}

// ImageTagCreateSummary represents the ImageTagCreateSummary schema from the OpenAPI specification
type ImageTagCreateSummary struct {
	Exceeded []ImageTagCreateEntry `json:"exceeded,omitempty"`
	Created []ImageTagCreateEntry `json:"created,omitempty"`
	Duplicated []ImageTagCreateEntry `json:"duplicated,omitempty"`
}

// Prediction represents the Prediction schema from the OpenAPI specification
type Prediction struct {
	Boundingbox BoundingBox `json:"boundingBox,omitempty"` // Bounding box that defines a region of an image.
	Probability float32 `json:"probability,omitempty"` // Probability of the tag.
	Tagid string `json:"tagId,omitempty"` // Id of the predicted tag.
	Tagname string `json:"tagName,omitempty"` // Name of the predicted tag.
}

// Region represents the Region schema from the OpenAPI specification
type Region struct {
	Width float32 `json:"width"` // Width.
	Height float32 `json:"height"` // Height.
	Left float32 `json:"left"` // Coordinate of the left boundary.
	Tagid string `json:"tagId"` // Id of the tag associated with this region.
	Top float32 `json:"top"` // Coordinate of the top boundary.
}

// SuggestedTagAndRegion represents the SuggestedTagAndRegion schema from the OpenAPI specification
type SuggestedTagAndRegion struct {
	Created string `json:"created,omitempty"` // Date this prediction was created.
	Id string `json:"id,omitempty"` // Prediction Id.
	Iteration string `json:"iteration,omitempty"` // Iteration Id.
	Predictionuncertainty float64 `json:"predictionUncertainty,omitempty"` // Uncertainty (entropy) of suggested tags or regions per image.
	Predictions []Prediction `json:"predictions,omitempty"` // List of predictions.
	Project string `json:"project,omitempty"` // Project Id.
}

// ImageCreateSummary represents the ImageCreateSummary schema from the OpenAPI specification
type ImageCreateSummary struct {
	Images []ImageCreateResult `json:"images,omitempty"` // List of the image creation results.
	Isbatchsuccessful bool `json:"isBatchSuccessful,omitempty"` // True if all of the images in the batch were created successfully, otherwise false.
}

// PredictionQueryTag represents the PredictionQueryTag schema from the OpenAPI specification
type PredictionQueryTag struct {
	Id string `json:"id,omitempty"`
	Maxthreshold float32 `json:"maxThreshold,omitempty"`
	Minthreshold float32 `json:"minThreshold,omitempty"`
}

// RegionProposal represents the RegionProposal schema from the OpenAPI specification
type RegionProposal struct {
	Boundingbox BoundingBox `json:"boundingBox,omitempty"` // Bounding box that defines a region of an image.
	Confidence float32 `json:"confidence,omitempty"`
}

// ProjectSettings represents the ProjectSettings schema from the OpenAPI specification
type ProjectSettings struct {
	Classificationtype string `json:"classificationType,omitempty"` // Gets or sets the classification type of the project.
	Detectionparameters string `json:"detectionParameters,omitempty"` // Detection parameters in use, if any.
	Domainid string `json:"domainId,omitempty"` // Gets or sets the id of the Domain to use with this project.
	Imageprocessingsettings ImageProcessingSettings `json:"imageProcessingSettings,omitempty"` // Represents image preprocessing settings used by image augmentation.
	Targetexportplatforms []string `json:"targetExportPlatforms,omitempty"` // A list of ExportPlatform that the trained model should be able to support.
	Usenegativeset bool `json:"useNegativeSet,omitempty"` // Indicates if negative set is being used.
}

// PredictionQueryToken represents the PredictionQueryToken schema from the OpenAPI specification
type PredictionQueryToken struct {
	Continuation string `json:"continuation,omitempty"`
	Endtime string `json:"endTime,omitempty"`
	Iterationid string `json:"iterationId,omitempty"`
	Maxcount int `json:"maxCount,omitempty"`
	Orderby string `json:"orderBy,omitempty"`
	Starttime string `json:"startTime,omitempty"`
	Tags []PredictionQueryTag `json:"tags,omitempty"`
	Application string `json:"application,omitempty"`
	Session string `json:"session,omitempty"`
}

// TrainingParameters represents the TrainingParameters schema from the OpenAPI specification
type TrainingParameters struct {
	Selectedtags []string `json:"selectedTags,omitempty"` // List of tags selected for this training session, other tags in the project will be ignored.
}

// ImagePerformance represents the ImagePerformance schema from the OpenAPI specification
type ImagePerformance struct {
	Predictions []Prediction `json:"predictions,omitempty"`
	Width int `json:"width,omitempty"`
	Imageuri string `json:"imageUri,omitempty"`
	Tags []ImageTag `json:"tags,omitempty"`
	Created string `json:"created,omitempty"`
	Id string `json:"id,omitempty"`
	Thumbnailuri string `json:"thumbnailUri,omitempty"`
	Height int `json:"height,omitempty"`
	Regions []ImageRegion `json:"regions,omitempty"`
}

// ImageFileCreateEntry represents the ImageFileCreateEntry schema from the OpenAPI specification
type ImageFileCreateEntry struct {
	Regions []Region `json:"regions,omitempty"`
	Tagids []string `json:"tagIds,omitempty"`
	Contents string `json:"contents,omitempty"`
	Name string `json:"name,omitempty"`
}

// Project represents the Project schema from the OpenAPI specification
type Project struct {
	Name string `json:"name"` // Gets or sets the name of the project.
	Settings ProjectSettings `json:"settings"` // Represents settings associated with a project.
	Status string `json:"status,omitempty"` // Gets the status of the project.
	Thumbnailuri string `json:"thumbnailUri,omitempty"` // Gets the thumbnail url representing the image.
	Id string `json:"id,omitempty"` // Gets the project id.
	Created string `json:"created,omitempty"` // Gets the date this project was created.
	Description string `json:"description"` // Gets or sets the description of the project.
	Drmodeenabled bool `json:"drModeEnabled,omitempty"` // Gets if the Disaster Recovery (DR) mode is on, indicating the project is temporarily read-only.
	Lastmodified string `json:"lastModified,omitempty"` // Gets the date this project was last modified.
}

// TagPerformance represents the TagPerformance schema from the OpenAPI specification
type TagPerformance struct {
	Recallstddeviation float32 `json:"recallStdDeviation,omitempty"` // Gets the standard deviation for the recall.
	Averageprecision float32 `json:"averagePrecision,omitempty"` // Gets the average precision when applicable.
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Precision float32 `json:"precision,omitempty"` // Gets the precision.
	Precisionstddeviation float32 `json:"precisionStdDeviation,omitempty"` // Gets the standard deviation for the precision.
	Recall float32 `json:"recall,omitempty"` // Gets the recall.
}

// ImagePrediction represents the ImagePrediction schema from the OpenAPI specification
type ImagePrediction struct {
	Project string `json:"project,omitempty"` // Project Id.
	Created string `json:"created,omitempty"` // Date this prediction was created.
	Id string `json:"id,omitempty"` // Prediction Id.
	Iteration string `json:"iteration,omitempty"` // Iteration Id.
	Predictions []Prediction `json:"predictions,omitempty"` // List of predictions.
}

// StoredImagePrediction represents the StoredImagePrediction schema from the OpenAPI specification
type StoredImagePrediction struct {
	Resizedimageuri string `json:"resizedImageUri,omitempty"` // The URI to the (resized) prediction image.
	Id string `json:"id,omitempty"` // Prediction Id.
	Project string `json:"project,omitempty"` // Project Id.
	Iteration string `json:"iteration,omitempty"` // Iteration Id.
	Predictions []Prediction `json:"predictions,omitempty"` // List of predictions.
	Thumbnailuri string `json:"thumbnailUri,omitempty"` // The URI to the thumbnail of the original prediction image.
	Domain string `json:"domain,omitempty"` // Domain used for the prediction.
	Originalimageuri string `json:"originalImageUri,omitempty"` // The URI to the original prediction image.
	Created string `json:"created,omitempty"` // Date this prediction was created.
}

// Image represents the Image schema from the OpenAPI specification
type Image struct {
	Resizedimageuri string `json:"resizedImageUri,omitempty"` // The URI to the (resized) image used for training.
	Tags []ImageTag `json:"tags,omitempty"` // Tags associated with this image.
	Id string `json:"id,omitempty"` // Id of the image.
	Regions []ImageRegion `json:"regions,omitempty"` // Regions associated with this image.
	Width int `json:"width,omitempty"` // Width of the image.
	Created string `json:"created,omitempty"` // Date the image was created.
	Height int `json:"height,omitempty"` // Height of the image.
	Thumbnailuri string `json:"thumbnailUri,omitempty"` // The URI to the thumbnail of the original image.
	Originalimageuri string `json:"originalImageUri,omitempty"` // The URI to the original uploaded image.
}

// ImageRegion represents the ImageRegion schema from the OpenAPI specification
type ImageRegion struct {
	Created string `json:"created,omitempty"`
	Height float32 `json:"height"` // Height.
	Left float32 `json:"left"` // Coordinate of the left boundary.
	Regionid string `json:"regionId,omitempty"`
	Tagid string `json:"tagId"` // Id of the tag associated with this region.
	Tagname string `json:"tagName,omitempty"`
	Top float32 `json:"top"` // Coordinate of the top boundary.
	Width float32 `json:"width"` // Width.
}

// Iteration represents the Iteration schema from the OpenAPI specification
type Iteration struct {
	Trainedat string `json:"trainedAt,omitempty"` // Gets the time this iteration was last modified.
	Trainingtype string `json:"trainingType,omitempty"` // Gets the training type of the iteration.
	Name string `json:"name"` // Gets or sets the name of the iteration.
	Reservedbudgetinhours int `json:"reservedBudgetInHours,omitempty"` // Gets the reserved advanced training budget for the iteration.
	Lastmodified string `json:"lastModified,omitempty"` // Gets the time this iteration was last modified.
	Status string `json:"status,omitempty"` // Gets the current iteration status.
	Domainid string `json:"domainId,omitempty"` // Get or sets a guid of the domain the iteration has been trained on.
	Originalpublishresourceid string `json:"originalPublishResourceId,omitempty"` // Resource Provider Id this iteration was originally published to.
	Trainingtimeinminutes int `json:"trainingTimeInMinutes,omitempty"` // Gets the training time for the iteration.
	Created string `json:"created,omitempty"` // Gets the time this iteration was completed.
	Projectid string `json:"projectId,omitempty"` // Gets the project id of the iteration.
	Id string `json:"id,omitempty"` // Gets the id of the iteration.
	Classificationtype string `json:"classificationType,omitempty"` // Gets the classification type of the project.
	Exportable bool `json:"exportable,omitempty"` // Whether the iteration can be exported to another format for download.
	Exportableto []string `json:"exportableTo,omitempty"` // A set of platforms this iteration can export to.
	Publishname string `json:"publishName,omitempty"` // Name of the published model.
}

// StoredSuggestedTagAndRegion represents the StoredSuggestedTagAndRegion schema from the OpenAPI specification
type StoredSuggestedTagAndRegion struct {
	Id string `json:"id,omitempty"` // Prediction Id.
	Predictions []Prediction `json:"predictions,omitempty"` // List of predictions.
	Resizedimageuri string `json:"resizedImageUri,omitempty"` // The URI to the (resized) prediction image.
	Thumbnailuri string `json:"thumbnailUri,omitempty"` // The URI to the thumbnail of the original prediction image.
	Domain string `json:"domain,omitempty"` // Domain used for the prediction.
	Iteration string `json:"iteration,omitempty"` // Iteration Id.
	Originalimageuri string `json:"originalImageUri,omitempty"` // The URI to the original prediction image.
	Predictionuncertainty float64 `json:"predictionUncertainty,omitempty"` // Uncertainty (entropy) of suggested tags or regions per image.
	Project string `json:"project,omitempty"` // Project Id.
	Width int `json:"width,omitempty"` // Width of the resized image.
	Created string `json:"created,omitempty"` // Date this prediction was created.
	Height int `json:"height,omitempty"` // Height of the resized image.
}

// ImageRegionCreateSummary represents the ImageRegionCreateSummary schema from the OpenAPI specification
type ImageRegionCreateSummary struct {
	Exceeded []ImageRegionCreateEntry `json:"exceeded,omitempty"`
	Created []ImageRegionCreateResult `json:"created,omitempty"`
	Duplicated []ImageRegionCreateEntry `json:"duplicated,omitempty"`
}

// ImageUrlCreateBatch represents the ImageUrlCreateBatch schema from the OpenAPI specification
type ImageUrlCreateBatch struct {
	Tagids []string `json:"tagIds,omitempty"`
	Images []ImageUrlCreateEntry `json:"images,omitempty"`
}

// PredictionQueryResult represents the PredictionQueryResult schema from the OpenAPI specification
type PredictionQueryResult struct {
	Results []StoredImagePrediction `json:"results,omitempty"` // Result of an prediction request.
	Token PredictionQueryToken `json:"token,omitempty"`
}

// ImageRegionProposal represents the ImageRegionProposal schema from the OpenAPI specification
type ImageRegionProposal struct {
	Projectid string `json:"projectId,omitempty"`
	Proposals []RegionProposal `json:"proposals,omitempty"`
	Imageid string `json:"imageId,omitempty"`
}

// IterationPerformance represents the IterationPerformance schema from the OpenAPI specification
type IterationPerformance struct {
	Pertagperformance []TagPerformance `json:"perTagPerformance,omitempty"` // Gets the per-tag performance details for this iteration.
	Precision float32 `json:"precision,omitempty"` // Gets the precision.
	Precisionstddeviation float32 `json:"precisionStdDeviation,omitempty"` // Gets the standard deviation for the precision.
	Recall float32 `json:"recall,omitempty"` // Gets the recall.
	Recallstddeviation float32 `json:"recallStdDeviation,omitempty"` // Gets the standard deviation for the recall.
	Averageprecision float32 `json:"averagePrecision,omitempty"` // Gets the average precision when applicable.
}

// ImageIdCreateBatch represents the ImageIdCreateBatch schema from the OpenAPI specification
type ImageIdCreateBatch struct {
	Images []ImageIdCreateEntry `json:"images,omitempty"`
	Tagids []string `json:"tagIds,omitempty"`
}

// ImageRegionCreateResult represents the ImageRegionCreateResult schema from the OpenAPI specification
type ImageRegionCreateResult struct {
	Width float32 `json:"width"` // Width.
	Top float32 `json:"top"` // Coordinate of the top boundary.
	Tagid string `json:"tagId"` // Id of the tag associated with this region.
	Tagname string `json:"tagName,omitempty"`
	Imageid string `json:"imageId,omitempty"`
	Left float32 `json:"left"` // Coordinate of the left boundary.
	Regionid string `json:"regionId,omitempty"`
	Created string `json:"created,omitempty"`
	Height float32 `json:"height"` // Height.
}

// ImageTagCreateBatch represents the ImageTagCreateBatch schema from the OpenAPI specification
type ImageTagCreateBatch struct {
	Tags []ImageTagCreateEntry `json:"tags,omitempty"` // Image Tag entries to include in this batch.
}

// Domain represents the Domain schema from the OpenAPI specification
type Domain struct {
	Name string `json:"name,omitempty"`
	TypeField string `json:"type,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Exportable bool `json:"exportable,omitempty"`
	Id string `json:"id,omitempty"`
}

// ImageCreateResult represents the ImageCreateResult schema from the OpenAPI specification
type ImageCreateResult struct {
	Image Image `json:"image,omitempty"` // Image model to be sent as JSON.
	Sourceurl string `json:"sourceUrl,omitempty"` // Source URL of the image.
	Status string `json:"status,omitempty"` // Status of the image creation.
}

// ProjectExport represents the ProjectExport schema from the OpenAPI specification
type ProjectExport struct {
	Imagecount int `json:"imageCount,omitempty"` // Count of images that will be exported.
	Iterationcount int `json:"iterationCount,omitempty"` // Count of iterations that will be exported.
	Regioncount int `json:"regionCount,omitempty"` // Count of regions that will be exported.
	Tagcount int `json:"tagCount,omitempty"` // Count of tags that will be exported.
	Token string `json:"token,omitempty"` // Opaque token that should be passed to ImportProject to perform the import. This token grants access to import this project to all that have the token.
	Estimatedimporttimeinms int `json:"estimatedImportTimeInMS,omitempty"` // Estimated time this project will take to import, can change based on network connectivity and load between source and destination regions.
}

// SuggestedTagAndRegionQueryToken represents the SuggestedTagAndRegionQueryToken schema from the OpenAPI specification
type SuggestedTagAndRegionQueryToken struct {
	Threshold float64 `json:"threshold,omitempty"` // Confidence threshold to filter suggested tags on.
	Continuation string `json:"continuation,omitempty"` // Continuation Id for database pagination. Initially null but later used to paginate.
	Maxcount int `json:"maxCount,omitempty"` // Maximum number of results you want to be returned in the response.
	Session string `json:"session,omitempty"` // SessionId for database query. Initially set to null but later used to paginate.
	Sortby string `json:"sortBy,omitempty"` // OrderBy. Ordering mechanism for your results.
	Tagids []string `json:"tagIds,omitempty"` // Existing TagIds in project to filter suggested tags on.
}

// ImageProcessingSettings represents the ImageProcessingSettings schema from the OpenAPI specification
type ImageProcessingSettings struct {
	Augmentationmethods map[string]interface{} `json:"augmentationMethods,omitempty"` // Gets or sets enabled image transforms. The key corresponds to the transform name. If value is set to true, then correspondent transform is enabled. Otherwise this transform will not be used. Augmentation will be uniformly distributed among enabled transforms.
}

// ImageUrlCreateEntry represents the ImageUrlCreateEntry schema from the OpenAPI specification
type ImageUrlCreateEntry struct {
	Regions []Region `json:"regions,omitempty"`
	Tagids []string `json:"tagIds,omitempty"`
	Url string `json:"url"` // Url of the image.
}

// ImageRegionCreateEntry represents the ImageRegionCreateEntry schema from the OpenAPI specification
type ImageRegionCreateEntry struct {
	Height float32 `json:"height"` // Height.
	Imageid string `json:"imageId"` // Id of the image.
	Left float32 `json:"left"` // Coordinate of the left boundary.
	Tagid string `json:"tagId"` // Id of the tag associated with this region.
	Top float32 `json:"top"` // Coordinate of the top boundary.
	Width float32 `json:"width"` // Width.
}

// ImageTag represents the ImageTag schema from the OpenAPI specification
type ImageTag struct {
	Created string `json:"created,omitempty"`
	Tagid string `json:"tagId,omitempty"`
	Tagname string `json:"tagName,omitempty"`
}

// ImageTagCreateEntry represents the ImageTagCreateEntry schema from the OpenAPI specification
type ImageTagCreateEntry struct {
	Imageid string `json:"imageId,omitempty"` // Id of the image.
	Tagid string `json:"tagId,omitempty"` // Id of the tag.
}

// CustomVisionError represents the CustomVisionError schema from the OpenAPI specification
type CustomVisionError struct {
	Message string `json:"message"` // A message explaining the error reported by the service.
	Code string `json:"code"` // The error code.
}

// ImageFileCreateBatch represents the ImageFileCreateBatch schema from the OpenAPI specification
type ImageFileCreateBatch struct {
	Tagids []string `json:"tagIds,omitempty"`
	Images []ImageFileCreateEntry `json:"images,omitempty"`
}

// ImageRegionCreateBatch represents the ImageRegionCreateBatch schema from the OpenAPI specification
type ImageRegionCreateBatch struct {
	Regions []ImageRegionCreateEntry `json:"regions,omitempty"`
}

// TagFilter represents the TagFilter schema from the OpenAPI specification
type TagFilter struct {
	Threshold float64 `json:"threshold,omitempty"` // Confidence threshold to filter suggested tags on.
	Tagids []string `json:"tagIds,omitempty"` // Existing TagIds in project to get suggested tags count for.
}

// ImageIdCreateEntry represents the ImageIdCreateEntry schema from the OpenAPI specification
type ImageIdCreateEntry struct {
	Id string `json:"id,omitempty"` // Id of the image.
	Regions []Region `json:"regions,omitempty"`
	Tagids []string `json:"tagIds,omitempty"`
}
