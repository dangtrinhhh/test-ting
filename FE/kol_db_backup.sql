PGDMP  4    -            	    |            weallnet    16.1    16.1     �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    254503    weallnet    DATABASE     �   CREATE DATABASE weallnet WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_United States.1252';
    DROP DATABASE weallnet;
                postgres    false            �            1259    254524    KOL    TABLE       CREATE TABLE public."KOL" (
    "KolID" bigint NOT NULL,
    "UserProfileID" bigint,
    "Language" text,
    "Education" text,
    "ExpectedSalary" bigint,
    "ExpectedSalaryEnable" boolean,
    "ChannelSettingTypeID" bigint,
    "IDFrontURL" text,
    "IDBackURL" text,
    "PortraitURL" text,
    "RewardID" bigint,
    "PaymentMethodID" bigint,
    "TestimonialsID" bigint,
    "VerificationStatus" boolean,
    "Enabled" boolean,
    "ActiveDate" timestamp with time zone,
    "Active" boolean,
    "CreatedBy" text,
    "CreatedDate" timestamp with time zone,
    "ModifiedBy" text,
    "ModifiedDate" timestamp with time zone,
    "IsRemove" boolean,
    "IsOnBoarding" boolean,
    "Code" text,
    "PortraitRightURL" text,
    "PortraitLeftURL" text,
    "LivenessStatus" boolean
);
    DROP TABLE public."KOL";
       public         heap    postgres    false            �            1259    254523    KOL_KolID_seq    SEQUENCE     x   CREATE SEQUENCE public."KOL_KolID_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 &   DROP SEQUENCE public."KOL_KolID_seq";
       public          postgres    false    216            �           0    0    KOL_KolID_seq    SEQUENCE OWNED BY     E   ALTER SEQUENCE public."KOL_KolID_seq" OWNED BY public."KOL"."KolID";
          public          postgres    false    215            O           2604    254527 	   KOL KolID    DEFAULT     l   ALTER TABLE ONLY public."KOL" ALTER COLUMN "KolID" SET DEFAULT nextval('public."KOL_KolID_seq"'::regclass);
 <   ALTER TABLE public."KOL" ALTER COLUMN "KolID" DROP DEFAULT;
       public          postgres    false    215    216    216            �          0    254524    KOL 
   TABLE DATA           �  COPY public."KOL" ("KolID", "UserProfileID", "Language", "Education", "ExpectedSalary", "ExpectedSalaryEnable", "ChannelSettingTypeID", "IDFrontURL", "IDBackURL", "PortraitURL", "RewardID", "PaymentMethodID", "TestimonialsID", "VerificationStatus", "Enabled", "ActiveDate", "Active", "CreatedBy", "CreatedDate", "ModifiedBy", "ModifiedDate", "IsRemove", "IsOnBoarding", "Code", "PortraitRightURL", "PortraitLeftURL", "LivenessStatus") FROM stdin;
    public          postgres    false    216   }       �           0    0    KOL_KolID_seq    SEQUENCE SET     >   SELECT pg_catalog.setval('public."KOL_KolID_seq"', 1, false);
          public          postgres    false    215            Q           2606    254531    KOL KOL_pkey 
   CONSTRAINT     S   ALTER TABLE ONLY public."KOL"
    ADD CONSTRAINT "KOL_pkey" PRIMARY KEY ("KolID");
 :   ALTER TABLE ONLY public."KOL" DROP CONSTRAINT "KOL_pkey";
       public            postgres    false    216            �   K  x��MO�8@��W�m�&�r�n��a���ݕ���L�ݝ�n%iV���Z��A)��HA⋕����\~��4X�.�bkwU��Y�n�[��;��^�U��V%!��V�ڶ�]s{�*�e�r�{<��ժr��{4mU�8�_sW�m�0D ���NBn���@w,/�jy��x���n&�tM���2t��K�F�0]����"��s��D�ڗΛ6�yڟ�p��?N���߿��,�۔�՞A����.~���_ۖnC�d��
'��Cs��R��@�`��J�A���w�^�5q��;�l�����G�.M�l�,�p(ʦ�󶬜J�r���{͠&51��ܧe�� ��i�羈�ԺV��+���W8��P���P�|�"�HBM��M:Z��aԌZ�N���>4eѨ��E�}�Qg�:�Y#Sg$��"k&-.�1S��yE�*o�cW�����b/&�b?��2x���Ow���.a�M��������W�{�S�K����K|���D"/�'/�t�����K���yOwYV�-���U��us���K����K�B��ԥuI�M'-.�S����=���=w�O�w�Ί�e�]F�e~��d�2���?�e��e;�=i���y(���.�����Q#��v|�+`����e����r�;v4��>g�����W���>�� �@��do���~o�U�QւdoAC�����\����\o?uH�!Q�[��H���L;�>�Zа�����ܵ������Cp��;M���dS���~S�^�vl*hxs��g��j��l.�'`va�[��H���f-FÎ���z��������!��g|�dq�����x�`�b4�X\�@���;��ˆ?8j�9�I���]L����N� y�z�W�����h�;�4�il�w�h��w~k�Λ��!L ��
$���
S�%?V4t?��<��~o.�y)��y�O��1@�����Lc�Fk^�[�S�s�qtA�t�� �a��0��a�`c,��q�+�����`8��� E3�ˠ�3P�(���8��	�����~����     