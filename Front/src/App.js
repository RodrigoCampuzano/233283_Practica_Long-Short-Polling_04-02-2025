import React from 'react';
import AddProductForm from './components/AddProductForm';
import DiscountCounter from './components/DiscountCounter';
import LastAddedProduct from './components/LastAddedProduct';
import './styles.css';

function App() {
  return (
    <div className="app-container">
      <div className="content">
        <div className="form-section">
          <AddProductForm />
          <DiscountCounter />
        </div>
        <div className="product-section">
          <LastAddedProduct />
        </div>
      </div>
    </div>
  );
}

export default App;