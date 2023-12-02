import streamlit as st
import os
from src.image_converter import webp_to_jpg
from src.video_converter import media_to_mp4


def webp_to_jpg_app():
    in_image = st.file_uploader("Upload WebP File", type=["webp"])

    if st.button("Convert Image"):
        # save the image to a temporary file
        print(in_image)
        in_name = in_image.name

        tmp_name = f"/tmp/{in_name}"

        with open(tmp_name, "wb") as f:
            f.write(in_image.read())

        out_image = webp_to_jpg(tmp_name)
        st.image(out_image, caption="Output Image", use_column_width=True)


def media_to_mp4_app():
    in_video = st.file_uploader("Upload WebM, WebP, or GIF File", type=["webm", "webp", "gif"])

    if st.button("Convert Video"):
        progress_bar = st.progress(0, text="Converting...")
        in_name = in_video.name
        tmp_name = f"/tmp/{in_name}"
        with open(tmp_name, "wb") as f:
            f.write(in_video.read())

        out_video = media_to_mp4(tmp_name, progress_bar)
        st.video(out_video)


def main():
    with st.expander(label="webp to jpg", expanded=True):
        st.header("webp to jpg")

        webp_to_jpg_app()

    with st.expander(label="media to mp4"):
        st.header("media to mp4")
        media_to_mp4_app()


if __name__ == "__main__":
    main()
