from pywebio import start_server
from pywebio.input import file_upload, input, select
from pywebio.output import put_image, put_video
from src.image_converter import webp_to_jpg
from src.video_converter import media_to_mp4


def webp_to_jpg_app():
    put_image('', src='')
    in_image = file_upload("Upload WebP File", accept=".webp")
    put_image('', src=in_image['content'])

    if input("Click 'Convert' to start conversion", type=BUTTONS) == 'Convert':
        out_image = webp_to_jpg(in_image['content'])
        put_image('', src=out_image)


def media_to_mp4_app():
    put_video('', src='')
    in_video = file_upload("Upload WebM, WebP, or GIF File", accept=".webm, .webp, .gif")
    put_video('', src=in_video['content'])

    if input("Click 'Convert' to start conversion", type=BUTTONS) == 'Convert':
        out_video = media_to_mp4(in_video['content'])
        put_video('', src=out_video)


def main():
    start_server({
        "webp to jpg": webp_to_jpg_app,
        "to mp4": media_to_mp4_app,
    }, debug=True, port=8088)


if __name__ == '__main__':
    main()
