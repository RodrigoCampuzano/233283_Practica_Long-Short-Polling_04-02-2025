import React, { useState, useEffect } from 'react';
import axios from 'axios';

const DiscountCounter = () => {
  const [count, setCount] = useState(0);
  const [isLoading, setIsLoading] = useState(true);

  const fetchDiscountCount = async () => {
    try {
      const response = await axios.get('http://localhost:8080/CountProductIsInDiscount');
      setCount(response.data.count);
    } catch (error) {
      console.error('Error fetching discount count:', error);
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    // Initial fetch
    fetchDiscountCount();

    // Long polling implementation
    const poll = () => {
      fetchDiscountCount();
      setTimeout(poll, 5000); // Poll every 5 seconds
    };

    const timeoutId = setTimeout(poll, 5000);

    return () => {
      clearTimeout(timeoutId);
    };
  }, []);

  return (
    <div className="discount-counter">
      <h2>Products on Discount</h2>
      {isLoading ? (
        <p>Loading...</p>
      ) : (
        <div className="count">{count}</div>
      )}
    </div>
  );
};

export default DiscountCounter;