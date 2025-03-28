import React, { useState, useEffect } from 'react';
import axios from 'axios';

const LastAddedProduct = () => {
  const [products, setProducts] = useState([]);
  const [isLoading, setIsLoading] = useState(true);

  const fetchRecentProducts = async () => {
    try {
      const response = await axios.get('http://localhost:8080/isNewProductAdded');
      if (response.data.products && response.data.products.length > 0) {
        setProducts(response.data.products);
      }
    } catch (error) {
      console.error('Error fetching recent products:', error);
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchRecentProducts();
    const interval = setInterval(fetchRecentProducts, 2000);
    return () => clearInterval(interval);
  }, []);

  return (
    <div className="last-product">
      <h2>Recently Added Products</h2>
      {isLoading ? (
        <p>Loading products...</p>
      ) : products.length === 0 ? (
        <p>No products added yet</p>
      ) : (
        <div className="products-list">
          {products.map((product, index) => (
            <div key={`${product.codigo}-${index}`} className="product-card">
              <h3>{product.nombre}</h3>
              <p><strong>Price:</strong> ${product.precio}</p>
              <p><strong>Code:</strong> {product.codigo}</p>
              <p className={`discount-status ${product.descuento ? 'on-discount' : ''}`}>
                {product.descuento ? 'On Discount' : 'Regular Price'}
              </p>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default LastAddedProduct;