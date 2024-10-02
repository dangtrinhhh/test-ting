import React from 'react';
import { Carousel } from 'antd';

const contentStyle: React.CSSProperties = {
  margin: 0,
  height: '160px',
  color: '#fff',
  lineHeight: '160px',
  textAlign: 'center',
  background: '#364d79',
};

// Define data type for Kols
interface Kols {
    kolID: number;
    userProfileID: number;
    language: string;
    education: string;
    expectedSalary: number;
    expectedSalaryEnable: boolean;
    channelSettingTypeID: number;
    iDFrontURL: string;
    iDBackURL: string;
    portraitURL: string;
    rewardID: number;
    paymentMethodID: number;
    testimonialsID: number;
    verificationStatus: boolean;
    enabled: boolean;
    activeDate: Date;
    active: boolean;
    createdBy: string;
    createdDate: Date;
    modifiedBy: string;
    modifiedDate: Date;
    isRemove: boolean;
    isOnBoarding: boolean;
    code: string;
    portraitRightURL: string;
    portraitLeftURL: string;
    livenessStatus: boolean;
}

interface MyCarouselProps {
    kols: Kols[];
}

const MyCarousel: React.FC<MyCarouselProps> = ({ kols }) => {

    // If kols is not an array, return early to avoid the error
    if (!Array.isArray(kols)) {
        return <p>No data available</p>;
    }

    return (
        <Carousel className="kols-carousel" arrows infinite={false}>
            { kols && kols.map((kol) => (
                <div key={kol.kolID} className="carousel-item">
                    <img src={kol.portraitURL} alt={kol.language} className="kol-image" />
                    <h2 style={contentStyle}>{kol.education}</h2>
                </div>
            ))}
        </Carousel>
    );
};

export default MyCarousel;
