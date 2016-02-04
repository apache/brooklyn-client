package models

type VersionSummary struct {
    Version     string                   `json:"version"`
    BuildSha1   string                   `json:"buildSha1"`
    BuildBranch string                   `json:"buildBranch"`
    Features    []BrooklynFeatureSummary `json:"features"`
}

type BrooklynFeatureSummary struct {
    Name           string            `json:"name"`
    SymbolicName   string            `json:"symbolicName"`
    Version        string            `json:"version"`
    LastModified   string            `json:"lastModified"`
    AdditionalData map[string]string `json:"additionalData"`
}

