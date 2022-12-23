package dto


type TreeAuthBo struct{
    NodePath  string `json:"nodePath"`
    NodeName  string `json:"nodeName"`
    NodeActions  []string `json:"nodeActions"`
    NodeValue  string `json:"nodeValue,omitempty"`
}

