package entities

type Product struct {
    ID        int    `json:"id"`
    Nombre    string `json:"nombre"`
    Precio    string `json:"precio"`
    Codigo    string `json:"codigo"`
    Descuento bool   `json:"descuento"`
}