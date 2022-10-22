package dto


type OidcDiscoveryMetadata struct{
    Issuer  string `json:"issuer"`
    AuthorizationEndpoint  string `json:"authorization_endpoint"`
    TokenEndpoint  string `json:"token_endpoint"`
    UserinfoEndpoint  string `json:"userinfo_endpoint"`
    JwksUri  string `json:"jwks_uri"`
    ScopesSupported  []string `json:"scopes_supported"`
    ResponseTypesSupported  []string `json:"response_types_supported"`
    ResponseModesSupported  []string `json:"response_modes_supported"`
    GrantTypesSupported  []string `json:"grant_types_supported"`
    IdTokenSigningAlgValuesSupported  []string `json:"id_token_signing_alg_values_supported"`
    IdTokenEncryptionAlgValuesSupported  []string `json:"id_token_encryption_alg_values_supported"`
    IdTokenEncryptionEncValuesSupported  []string `json:"id_token_encryption_enc_values_supported"`
    UserinfoSigningAlgValuesSupported  []string `json:"userinfo_signing_alg_values_supported"`
    UserinfoEncryptionAlgValuesSupported  []string `json:"userinfo_encryption_alg_values_supported"`
    UserinfoEncryptionEncValuesSupported  []string `json:"userinfo_encryption_enc_values_supported"`
    TokenEndpointAuthMethodsSupported  []string `json:"token_endpoint_auth_methods_supported"`
    ClaimTypesSupported  []string `json:"claim_types_supported"`
    ClaimsSupported  []string `json:"claims_supported"`
    CodeChallengeMethodsSupported  []string `json:"code_challenge_methods_supported"`
    EndSessionEndpoint  string `json:"end_session_endpoint"`
    IntrospectionEndpoint  string `json:"introspection_endpoint"`
    IntrospectionEndpointAuthMethodsSupported  []string `json:"introspection_endpoint_auth_methods_supported"`
    RevocationEndpoint  string `json:"revocation_endpoint"`
    RevocationEndpointAuthMethodsSupported  []string `json:"revocation_endpoint_auth_methods_supported"`
}

