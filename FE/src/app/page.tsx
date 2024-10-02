"use client"

import React, { useEffect, useState } from 'react';
import MyCarousel from '@/components/MyCarousel';
import axiosInstance from '@/api/axios';

// * Import CSS file, you can use CSS module if you want
// ! Change your CSS inside this file
import './page.css'

interface Kols {
	KolID: number;
	UserProfileID: number;
	Language: string;
	Education: string;
	ExpectedSalary: number;
	ExpectedSalaryEnable: boolean;
	ChannelSettingTypeID: number;
	IDFrontURL: string;
	IDBackURL: string;
	PortraitURL: string;
	RewardID: number;
	PaymentMethodID: number;
	TestimonialsID: number;
	VerificationStatus: boolean;
	Enabled: boolean;
	ActiveDate: Date;
	Active: boolean;
	CreatedBy: string;
	CreatedDate: Date;
	ModifiedBy: string;
	ModifiedDate: Date;
	IsRemove: boolean;
	IsOnBoarding: boolean;
	Code: string;
	PortraitRightURL: string;
	PortraitLeftURL: string;
	LivenessStatus: boolean;
}

const Page = () => {
	const [Kols, setKols] = useState<Kols[]>([]);

	const fetchKols = async () => {
		try {
			const response = await axiosInstance.get<Kols[]>('/kols?pageIndex=1&pageSize=20');

			// Log to check response data
			console.log("🚀 ~ response.data:", response.data)
			if (response.data?.kol) {
				setKols(response.data.kol);
			}
		} catch (err: any) {
			console.error(err);
		}
	};

	useEffect(() => {
		fetchKols();
	}, []);

	return (
		<>
			<h1 className='header'>Let's check our well-known KOLS!</h1>
			<MyCarousel kols={Kols} className='carousel-section' />
		</>
	);
};

export default Page;