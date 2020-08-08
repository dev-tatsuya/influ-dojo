import 'package:json_annotation/json_annotation.dart';

part 'rank_user.g.dart';

@JsonSerializable(fieldRename: FieldRename.snake)
class RankUser {
  RankUser(
      this.name, this.screenName, this.profileImage, this.point, this.ranking, this.lastRanking);

  String name;
  String screenName;
  String profileImage;
  num point;
  int ranking;
  int lastRanking;

  factory RankUser.fromJson(Map<String, dynamic> json) => _$RankUserFromJson(json);
  Map<String, dynamic> toJson() => _$RankUserToJson(this);
}
