""" """

import string
from pathlib import Path

import requests
from shapely.geometry import shape

districts = {
    "H1": "LRH",
    "H2": "LRL",
    "H3": "LRN",
    "H4": "LRP",
    "H5": "LRB",
    "H6": "LRC",
    "H7": "LRE",
    "B1": "MVM",
    "B2": "MVN",
    "B3": "MVS",
    "B4": "MVK",
    "B5": "MVR",
    "B6": "MVP",
    "E1": "NAB",
    "E3": "NAN",
    "E4": "NAO",
    "E5": "NAP",
    "E6": "NAE",
    "E7": "NAU",
    "G2": "NWP",
    "G3": "NWS",
    "G4": "NWW",
    "G5": "NWK",
    "G6": "NWO",
    "J1": "POF",
    "J2": "POJ",
    "J3": "POH",
    "J4": "POA",
    "K2": "SAC",
    "K3": "SAJ",
    "K5": "SAM",
    "K6": "SAS",
    "K7": "SAW",
    "K8": "SAA",
    "L1": "SPL",
    "L2": "SPK",
    "L3": "SPN",
    "L4": "SPA",
    "M2": "SWF",
    "M3": "SWG",
    "M4": "SWL",
    "M5": "SWT",
    # "N2": "TED",
    # "P0": "GRD",
    # "E2": "WAD",
}

# missions dictonary with the key making up the suffix of the table name is goes to
# e.g., office_aor_cw
missions = {
    "cw": string.Template(
        "https://services7.arcgis.com/n1YM8pTrFmm7L4hs/arcgis/rest/services/usace_cw_districts/FeatureServer/0/query?outFields=*&where=SYMBOL%3D%27${SYMBOL}%27&f=geojson"
    ),
    "reg": string.Template(
        "https://services7.arcgis.com/n1YM8pTrFmm7L4hs/arcgis/rest/services/usace_regulatory_boundary/FeatureServer/0/query?outFields=*&where=DIST_ABBR%3D%27${SYMBOL}%27&f=geojson"
    ),
    "fuds": string.Template(
        "https://services7.arcgis.com/n1YM8pTrFmm7L4hs/arcgis/rest/services/fuds/FeatureServer/9/query?outFields=*&where=DIST%3D%27${SYMBOL}%27&f=geojson"
    ),
    "mil": string.Template(
        "https://services7.arcgis.com/n1YM8pTrFmm7L4hs/arcgis/rest/services/usace_mil_dist/FeatureServer/0/query?outFields=*&where=DIST%3D%27${SYMBOL}%27&f=geojson"
    ),
}

schema = "usace"

parent = Path(__file__).parent
dist_geom = parent / "dist_geom"
# geom_zip = parent / "mission_geom.zip"

# with ZipFile(geom_zip, mode="w") as zipp:
for dist_id, dist in districts.items():
    print(dist_id, dist)
    filename = dist_geom / f"{dist}_mission_geom.sql"

    write_lines = []
    for mission, url in missions.items():
        dbtable = f"office_aor_{mission}"
        url = url.substitute(SYMBOL=dist)
        resp = requests.get(url)
        resp_json = resp.json()

        features = resp_json.get("features")

        if len(features) > 0:
            write_lines.append(
                f"INSERT INTO {schema}.{dbtable} (office_id, geom) VALUES "
            )
            feature = features[0]
            geometry = feature.get("geometry")
            geom_shape = shape(geometry)
            wkt = geom_shape.wkt

            write_lines.append(f"('{dist_id}', ST_GeomFromText('{wkt}', 4326))")
            write_lines.append(";\n")

        with filename.open("w") as fp:
            fp.writelines(write_lines)
        # zipp.writestr(f"{dist}_mission_geom.sql", "".join(write_lines))
